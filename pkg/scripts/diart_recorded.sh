#!/bin/bash
source /home/rahma/miniconda3/etc/profile.d/conda.sh
conda activate diart
python3 -m diart.stream /home/rahma/recorded_audio/ --output /home/rahma/sound_output

