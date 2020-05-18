provider dummy {
  directory = "/tmp"
}

resource dummy_file my_file {
  file_name = "file.tmp"
  contents  = "initial contents\n"
}
