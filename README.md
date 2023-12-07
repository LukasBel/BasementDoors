# Weather Alert Program
### Using an API Key for Tomorrow.io, I parse the json weather data and unmarshall it into a weather struct. The program is designed to run indefinitely with calls to the API every 30 minutes to check for precipitation probability. If the probability is 50% or greater, I get an email notifying me to close the basement doors.


### Written in Go.

### Version 0.1
- Notifies list of emails regarding the precipitation probability 🌧️

### Version 0.2
- Added 🌨️ snowfall intensity feature
