package main


/*

curl -XPUT localhost:8301/store/client/_mapping -d '
{
    "client" : {
        "properties" : {
            "local_ip" : {"type" : "string", "store" : "yes"}
        }
    }
}

 */
