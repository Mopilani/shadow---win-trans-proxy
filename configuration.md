# Config File

```
{
    // Proxy Server

    // Shadowsocks
    // ss://ciphername:password@ip:port

    // Trojan-(GFW/GO)
    // trojan://password@ip:port#domain.name
    // Trojan-GO
    // trojan://password@ip:port/
    //   path?
    //   transport=(tls|websocket)
    //   &mux=(off|v1)
    //   #domain.name

    // supported cipher name: CHACHA20-IETF-POLY1305, AES-256-GCM, DUMMY
    "server": "ss://chacha20-ietf-poly1305:password@127.0.0.1:8388",


    // DNS Server
    // tls://1.1.1.1
    // https://1.1.1.1/dns-query
    "name_server": "https://1.1.1.1/dns-query",


    // tun device only
    "tun_name": "utun",
    "tun_addr": ["192.168.0.11/24"],


    // windivert only
    // filter string passed to WinDivert
    // https://www.reqrypt.org/windivert-doc.html#filter_language
    // outbound and ip and ip.DstAddr != serverip and ip.DstAddr != dnsserverip
    // example: outbound and ip and ip.DstAddr != 1.2.3.4 and ip.DstAddr != 1.1.1.1
    "windivert_filter_string": "outbound and ip and ip.DstAddr != 1.2.3.4 and ip.DstAddr != 1.1.1.1",

    // windivert only
    // maxmind geoip file
    // proxy/bypass = iso code of country
    // final = proxy/bypass
    "geo_ip_rules": {
        "file": "",
        "proxy": [],
        "bypass": [],
        "final": "",
    },

    // windivert only
    // programs in this list will be proxied
    "app_rules": {
        "proxy":[
            "git.exe"
        ]
    },


    // Packets to IPs in this list will be diverted to shadow
    // For tun device, these IPs will be added to route table
    // For WinDivert, packets sending to these IPs will be diverted
    "ip_cidr_rules": {
        "proxy": [
            "198.18.0.0/16",
            "8.8.8.8/32"
        ]
    },


    // shadow will hijack all UDP dns queries
    // domains in proxy list will be given a fake ip: 198.18.X.Y
    // and drop all queries for domains in blocked
    // and redirect queries to name_server for domains in direct.
    // If not found, it is direct
    "domain_rules": {
        "proxy": [
            "**.google.com",
            "**.google.*",
            "**.google.*.*",
            "**.youtube.com",
            "*.twitter.com",
            "www.facebook.com",
            "bing.com",
            "**.amazon.*"
        ],
        "direct": [
            "**.baidu.*",
            "**.youku.*",
            "**.*"
        ],
        "blocked": [
            "ad.blocked.com"
        ]
    }
}
```
