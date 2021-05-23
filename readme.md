# Real Estate
This project is an effort to monitoring the real estate market prices in Greece.

The code fetches all available houses based on the queries given on main.go and pushes the data to a prometheus server, whose address is provided via an environment variable.

## Building & running
`make build run REPO=hub.docker.com TARGET=mac(default)|arm` will create an image for you. By default the name of the image contains my private docker repo but this can be overriden by using the REPO variable in the make command.

## Helm Chart
There is a helmchart available in the repository. One can install it by using `helm install real_estate helmchart/ -n real_estate -f your_values.yml`

## Next steps
Add more sites, due to google recaptcha protection it was not possible to fetch data from xe.gr, spitogatos.gr.
Facebook probably has similar issues.

### Dashboard
- [x] Create dropdown to select multiple regions
- [x] price/m2
- [ ] Show the first and last price on best deals
- [ ] Show the avg time an ad is online and the price difference

### TODO
- [x] provide any queries throught the environment
- [ ] write tests (due to the limited time i have right now i decided to start with something that would start saving data and then worry abou testing
- [ ] set up alerts on new ads
- [ ] provide grafana dashboard configmap 
- [ ] provide alerts config
- [ ] refactor utils into a struct and rename
- [ ] concurrent queries with chains
- [ ] add more sources
- [ ] create a pipeline on github

