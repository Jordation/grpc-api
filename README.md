# API 
Re-Architecting my golang graph generation/valorant stats API to utilise grpc, protobuffers to handle internal communications.
The main project and all the logic for the API lives [here](https://github.com/Jordation/go-api)

## Plans
Create grpc endpoints and API client/server so it's not all running through a struct/interface composed of the individual services
Merge this structure and new architechtural decisions back into the core app 