# kubepacket WIP

Count kubernetes pod packets by adding annotations

### packet count metric example

    packet_count{dst_ip="192.168.1.27",dst_port="8080",proto="tcp",src_ip="192.168.1.27",src_port="dynamic"} 176
    packet_count{dst_ip="192.168.1.27",dst_port="dynamic",proto="tcp",src_ip="192.168.1.27",src_port="8080"} 134

### kubernetes integration

Based controll on the excellent controller from https://github.com/bitnami-labs/kubewatch
