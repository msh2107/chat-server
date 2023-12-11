package service

//go:generate sh -c "rm -rf mocks && mkdir -p mocks && pwd"
//go:generate ../../bin/minimock -i ChatService -o ./mocks/ -s "_minimock.go"
