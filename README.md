

steps to execute:
1. cd scheduleclock
2. go mod init clock
3. go build
4. ./clock
 
 and check that the text in secsFile is displaying for every second, text in mntsFile per minute & text in hoursFile every hour to stdout.
 
5. while Running, user can change the text in files: secsFile/mntsFile/hoursFile, so that corresponding text will be displayed.

(OR)
steps to execute:
1. cd scheduleclock
2. go mod init clock
3. go build
4. go test -v -timeout 3h -coverprofile=cover.out && go tool cover -html=cover.out


Note: 
run the command below to get the coverage:
go tool cover -html=cover.out 

Please open the file : test_coverage.png  for coverage report.