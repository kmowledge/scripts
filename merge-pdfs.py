#! /usr/bin/python3

from pypdf import PdfMerger
import sys

# pdfs = []
# pdfs.append(sys.argv[1:])

merger = PdfMerger()

for arg in sys.argv[1:]:
      merger.append(arg)

merger.write('merged.pdf')
merger.close()

sys.exit()