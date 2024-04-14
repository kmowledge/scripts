#! /usr/bin/python3

import webbrowser, sys, time

arg = '' # without this->[xxxxx:xxxxx:0100/000000.xxxxxx:ERROR:zygote_linux.cc(672)] write: Broken pipe (32)

for arg in sys.argv[1:]:
      url = ['https://', arg, '/']
      webbrowser.open(''.join(url))

time.sleep(3)
sys.exit()