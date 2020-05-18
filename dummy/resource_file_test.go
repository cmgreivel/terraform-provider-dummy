package dummy

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func testAccFileDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "dummy_file" {
			continue
		}
		_, err := os.Stat(rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Stat for %s did not return error", rs.Primary.ID)
		}
		if ! os.IsNotExist(err) {
			return fmt.Errorf("Expexted NotExist(err), got %v", err)
		}
	}
	return nil
}


func openFile(s *terraform.State, r string) (*os.File, error) {
	rs, ok := s.RootModule().Resources[r]
	if !ok {
		return nil, fmt.Errorf("Not found: %s", r)
	}
	return os.Open(rs.Primary.ID)
}

func testAccCheckFileExists(r string, c string, t *string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		f, err := openFile(s, r)
		if err != nil {
			return fmt.Errorf("Could not open file for resource %s", r)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		contentsFound := false
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "CREATION: ") && t != nil {
				// Save the timestamp line
				*t = line
				log.Printf("[DEBUG] CMG: Saving %s\n", *t);
			}
			if strings.Contains(line, "CONTENTS: ") {
				contentsFound = true
				expected, err := formatContentsLine(c)
				if err != nil {
					return err
				}
				if line !=  expected {
					return fmt.Errorf("Expected '%s', got '%s'\n", expected, line)
				}
			}
		}
		if ! contentsFound {
			return fmt.Errorf("Did not find contents line in file for resource %s\n", r)
		}
		return nil
	}
}

func testAccCheckTimestamp(orig *string, new *string, match bool) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		if len(*new) == 0 || len(*orig) == 0 {
			return fmt.Errorf("At least one timestamp has zero length. orig '%s', new '%s'",
				*orig, *new)
		}
		if match {
			if *new != *orig {
				return fmt.Errorf("No match: '%s' != '%s'\n", *orig, *new)
			}
		} else {
			if *new == *orig {
				return fmt.Errorf("Unexpected match: '%s' != '%s'\n", *orig, *new)
			}
		}
		return nil
	}
}

const testAccGoodFileSpec = `
resource dummy_file test_file {
  file_name = "testfile.tmp"
  contents = "test contents"
}
`

func TestAccInitialCreation(t *testing.T) {
	resource.Test(t, resource.TestCase {
		PreCheck: func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: testAccFileDestroy,
		Steps: []resource.TestStep {
			{
				Config: testAccGoodFileSpec,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFileExists("dummy_file.test_file", "test contents", nil),
				),
			},
		},
	})
}

const testAccModifiedContentsSpec = `
resource dummy_file test_file {
  file_name = "testfile.tmp"
  contents = "modified contents"
}
`

const testAccModifiedFilenameSpec = `
resource dummy_file test_file {
  file_name = "new_testfile.tmp"
  contents = "modified contents"
}
`

func TestAccResourceUpdated(t *testing.T) {
	var origContent, newContent string

	resource.Test(t, resource.TestCase {
		PreCheck: func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		CheckDestroy: testAccFileDestroy,
		Steps: []resource.TestStep {
			{
				Config: testAccGoodFileSpec,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFileExists("dummy_file.test_file", "test contents", &origContent),
				),
			},
			{
				Config: testAccModifiedContentsSpec,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFileExists("dummy_file.test_file", "modified contents", &newContent),
					testAccCheckTimestamp(&origContent, &newContent, true),
				),
			},
			{
				Config: testAccModifiedFilenameSpec,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFileExists("dummy_file.test_file", "modified contents", &newContent),
					testAccCheckTimestamp(&origContent, &newContent, false),
				),
			},
		},
	})
}
