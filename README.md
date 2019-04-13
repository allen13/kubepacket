# kubepacket WIP
Count kubernetes pod packets. Enable using annotations. Kubernetes integration not yet started.

## packet count metric example

    packet_count{dst_ip="192.168.1.27",dst_port="8080",proto="tcp",src_ip="192.168.1.27",src_port="dynamic"} 176
    packet_count{dst_ip="192.168.1.27",dst_port="dynamic",proto="tcp",src_ip="192.168.1.27",src_port="8080"} 134
