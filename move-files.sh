#! /bin/bash

SRC="$1"
DEST="$2"

if [ "$#" -ne 2 ]; then
      echo "Provide exactly two arguments."
      exit 1
fi

if [ ! -d "$SRC" ]; then
      echo "Source path doesn't exist."
      exit 1
fi

if [ ! -d "$DEST" ]; then
      read -p "Destination path doesn't exist. Create? (y/n)" choice
      if [[ "$choice" == "y" ]]; then
            mkdir -p "$DEST"
      else
            exit 1
      fi
fi

find "$SRC" -mindepth 1 -maxdepth 1 -exec mv -t "$DEST" {} +
echo "Przeniesiono pliki i ścieżki z '$SRC' do '$DEST'."