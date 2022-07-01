#!/bin/bash
source /home/rahma/conda_init.sh
conda activate diart
python3 -m diart.stream microphone --output /home/rahma/out.txt