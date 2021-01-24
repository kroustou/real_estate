# Real Estate
This project is an effort to monitoring the real estate market prices in Greece.

The code fetches all available houses based on the queries given on main.go and pushes the data to a prometheus server, whose address is provided via an environment variable.

## Next steps
Add more sites, due to google recaptcha protection it was not possible to fetch data from xe.gr, spitogatos.gr.
Facebook probably has similar issues.

- [ ] provide any queries throught he environment
- [ ] set up alerts 
- [ ] provide grafana dashboard configmap 
- [ ] write tests (due to the limited time i have right now i decided to start with something that would create graphs and then worry about testing)
- [ ] add more sources

### Dashboard
- [ ] Create dropdown to select multiple regions
- [ ] Create price range selection
- [ ] Show the last price in ads that do not exist anymore
