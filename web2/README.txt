csvからsqliteへのimport
sqlite> .mode csv
sqlite> .import sys_patent.csv patent

または

$ sqlite3 -separator , db_file_path ".import sys_patent.csv patenr"
