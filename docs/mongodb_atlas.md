## Creating a free mongoDB connection

1. Create a free account at mongodb atlas: https://docs.atlas.mongodb.com/getting-started/
2. create and deply a new cluster : https://docs.atlas.mongodb.com/tutorial/deploy-free-tier-cluster/
3. Whitelist IP : https://docs.atlas.mongodb.com/tutorial/whitelist-connection-ip-address/
    1. Added IP address in home 
    2. Create a mongodb user and password
    3. Choose a connection method
        1. Connect your application : Go v 1.0+
        `mongodb+srv://pokke:<password>@cluster-1-osbty.gcp.mongodb.net/test?retryWrites=true&w=majority`
        2. Get an error 
        `error parsing uri: lookup _mongodb._tcp.cluster-1-osbty.gcp.mongodb.net on 192.168.2.1:53: no such host`
        NEED TO DEBUG this later!!!!
