#! /usr/bin/python3

import webbrowser, sys, time

for arg in sys.argv[1:]:
      url = ['https://', arg, '/']
      webbrowser.open(''.join(url))

time.sleep(3)
sys.exit()