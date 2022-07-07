#!/bin/bash
sleep 20
source /home/rahma/miniconda3/etc/profile.d/conda.sh
conda activate diart
python3 -m diart.stream microphone --output /home/rahma/sound_output

