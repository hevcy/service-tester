import schedule
import requests
import time
import os
import sys
from urllib3.exceptions import InsecureRequestWarning

# Attempts a get request on the provided address and prints the result.
def job(addr):
    t = time.localtime()
    req_time = time.strftime("%H:%M:%S", t)
    try:
        r = requests.get(addr, verify=False)
        print(f"{req_time} {addr} {str(r.status_code)}")
    except requests.exceptions.RequestException as ex :
        print(f"{req_time} {addr} REQUEST FAILED WITH: {ex}")

# This disables logging ssl warnings to avoid clogging up the console.
requests.packages.urllib3.disable_warnings(category=InsecureRequestWarning)

# Takes values in the env variable comma-separated, and splits these into a list to iterate over.
addresses = os.environ.get("SERVICE_TESTER_ENDPOINTS")
addrlist = addresses.split(",")
schedule_seconds = os.environ.get("SERVICE_TESTER_SCHEDULE_SECONDS")

# Adds jobs to the schedule for every address provided in the env variable at the defined frequency.
for x in addrlist:
    schedule.every(int(schedule_seconds)).seconds.do(job, addr=x)

# Checks for pending scheduled jobs every second. Runs indefinitely until killed.
while True:
    schedule.run_pending()
    time.sleep(1)