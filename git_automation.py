import os
import time

while True:
    os.system('git add .')
    os.system('git commit -m "hexa-arch"')
    os.system('git push -u origin main')
    time.sleep(60)  # sleep for 1 min
