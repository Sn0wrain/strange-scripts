#! /bin/bash

# Generate executable file by package name
# A solution for multi-tasks using one platform

GOSRC=$GOPATH/src
abs=`pwd`

# Check environment
if [ -d /usr/local/go1.7/ ]; then
	export GOROOT=/usr/local/go1.7
	export PATH=$GOROOT/bin:$PATH
fi

# Check package name 
if [[ ! -d $abs ]]; then
	echo "# $abs is not a dir."
	exit 3
fi
if [[ ! $abs =~ $GOSRC ]]; then
	echo "# $abs is not a go package."
	exit 7
fi

# Generate unique executable file name by package name
package=${abs##$GOSRC}
package=${package##/}
service=$package
service=${service%%/}
service=${service//\//_}

if [[ "$service" == "" ]]; then
	echo "# $abs is not a package."
	exit 4
fi
service="$service"

# TODO:Logs can be added here

go generate ./...
# Generate the build directory 
TmpDir="$abs/tmp"
mkdir $TmpDir
cd $TmpDir

echo "# Building $package ..."
go build -o "$service" $package || exit 6
if [ -x $service ]; then
	mv $service $abs/
fi
rm -rf $TmpDir
