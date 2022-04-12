# AR_WayFinding

###Github Repo:
https://github.com/WanghanlinStephen/AR_WayFinding
###API Documentation

`/v1/api/`
API used for AR Navigation Services
    
    search:
        Purpose: provide navigation direction( angle ) based on dijkstra algorithm.
    path:
        Purpose: provide shortest path information based on user destination and position
    connections:
        all:
            Purpose:return all conections as a list
        map:
            Purpose:require user input map id, then return all connections within this map
    nodes:
        all:
            Purpose:return all nodes as a list
        map:
            Purpose:require user input map id, then return all nodes within this map
        building:
            Purpose:require user input building name, then return all nodes within this building

`/v1/admin`
API used for Admin Management Services

    add:
        node:
            Purpose: add new node
        connection:
            Purpose: add new connection
        staircase:
            Purpose: update the status of a node to staircase
        map:
            Purpose: add new map
        
    delete:
        node:
            Purpose: delete new node
        connection:
            Purpose: delete new connection
        both:
            Purpose: delete connection and node at same time
        map:
            Purpose: delete new map
    
    index:
        nodeId:
            Purpose: return nodeId based on the information user input

    map:
        all:
            Purpose:return all maps
        name:
            Purpose:return the map based on the map name user input
        filter:
            id:
                Purpose:fetch map by id 
            name:
                Purpose:fetch map by name
        nodeId:
            Purpose:fetch map by node id
        building:
            Purpose:fetch building name by node id