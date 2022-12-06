
1. Ensure to add the following ports to your firewall:
- 2376
- 2377
- 7946
- 4789

2. Install Docker
`sudo apt install docker.io -y`

3. Initialize the cluster
`docker swarm init --advertise-addr node_ip_address`