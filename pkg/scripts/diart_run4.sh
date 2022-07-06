#!/bin/bash
source /home/rahma/miniconda3/etc/profile.d/conda.sh
conda activate diart
python3 -m diart.stream microphone --output sound_output

