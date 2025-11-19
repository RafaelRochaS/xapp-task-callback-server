# Callback Server for Edge Tasks
Small server to receive calls when services executed by edge devices are completed. 
The main idea is to accompany the [Edge Device Simulator](https://github.com/RafaelRochaS/edge-device-simulator), serving
as the callback server for tasks that are completed. 

Each request is written to a ```results.jsonl```, where request is a JSON line on the file. See [JSON Lines](https://jsonlines.org/)
for more information on the format. 

The repo includes a simple python script to analyze the output, and a simple Dockerfile to run the server.
Alternatively, the server can be run alongside the device simulators via the docker-compose file on the
[simulator repository](https://github.com/RafaelRochaS/edge-device-simulator/blob/main/scenario-0-compose.yml).
