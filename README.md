

steps to execute:
1. cd scheduleclock
2. go mod init clock
3. go build
4. ./clock
 check that the text in secsFile is displaying for every second, text in mntsFile per minute & text in hoursFile to stdout.
5. while Running, user can change the text in files: secsFile/mntsFile/hoursFile, so that corresponding text will be displayed.

(OR)
steps to execute:
1. cd scheduleclock
2. go mod init clock
3. go build
4. go test -v -timeout 3h -coverprofile=cover.out && go tool cover -html=cover.out
