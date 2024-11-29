#!/bin/bash

curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r '.[] | select(.id == 170) |.name , .powerstats.power, .appearance.gender'

#jq -r this option of -r it's help us to delete double quote " " or ouput it without " " 
