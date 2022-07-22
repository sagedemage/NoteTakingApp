#!/bin/bash

Help() {
    echo "-h	help"
    echo "-b	build the project"
	echo "-e	execute the project binary"
    echo "-r	build and run the project"
}

while getopts ":behr" arg; do
    
    case $arg in
        h) 
            Help
            exit;;
        b)
			go build -o build/out -v cmd/app/*.go
            exit;;
		e)
			./build/out
			exit;;
		r) 
			go run -v cmd/app/*.go
			exit;;
        \?)
            echo "Invalid option: pass -h for help"
            exit;;
    esac
done

if [ $OPTIND -eq 1 ]
then 
    echo "No arguments passed: pass -h for help"
fi
