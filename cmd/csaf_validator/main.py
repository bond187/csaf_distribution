import ctypes
import os

#test dir
directory = os.fsencode('./2023/')

#create object that refers to the .so file created 
lib = ctypes.CDLL('./main.so')
#run is now a callable function that is declared in main.go
run = lib.run

run.argtypes = [ctypes.c_char_p]

#loop over all 2023 alerts from 2023 test dir in same dir.
#notice that a few fail the check.
#run() will return 0 for pass, and 1 for fail. Can filter into directories based on this.
for _, _, files in os.walk(directory):
    for file in files:
        string = file.decode('utf-8')
        filename = "./2023/" + string
        run(filename.encode('utf-8'))
