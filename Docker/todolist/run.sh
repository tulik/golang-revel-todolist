#!/bin/sh
SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)
"$SCRIPTPATH/todolist" -importPath todolist -srcPath "$SCRIPTPATH/src" -runMode dev
