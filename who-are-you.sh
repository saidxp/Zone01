#!/bin/bash

curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq '.[] | select(.id == 70) | .name'

# curl it's tool who fetch file 
# -s i specific the all.json
# '.' it's for acc to element of file json
# []accessing by arrays
# | PIPE HELPD US TO TREAT ELEMENT ONE BY ONE
#SELECT() FUNCTION WHO RETURN ARRAY JSON WHO IN ID 70
#.name it the way who i spesific the element "n
