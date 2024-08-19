# Service Tester

This helm chart deploys a simple service that runs a python script to generate get requests.

## Configuration

This one is simple, there are only two config options! You need to pass in a list of endpoints you want to make a request to into the `SERVICE_TESTER_ENDPOINTS` environment variable, and then you can specify in `SERVICE_TESTER_SCHEDULE_SECONDS` the time between each attempt of polling all the endpoints in the list, in seconds of course.

## Job Script

In `/scripts`, you'll find job.py, which gets executed once the pod starts up. In the values file, you can see the sequence of commands that are run in a bash shell beforehand, this is just so we can make easy adjustments to if if needed. At present, I'm just pip installing two required dependencies, and then running the script. You can find more detail on what the script is actually doing in the comments on the script itself.