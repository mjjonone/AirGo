declare interface NodeInfoOld {
    created_at: string;
    updated_at: string;
    id: number;
    //基础参数
    remarks: string;
    address: string;
    port: string;
    node_order: number;        //节点排序
    enabled: boolean;

    sort: number;             //类型sort==11  V2Ray vmess
    nodespeed_limit: number;  //限速
    traffic_rate: number;     //倍率
    status: boolean;          //节点状态，true启用，flase禁用
    enable_transfer: boolean; //是否启用中转
    transfer_address: string; //中转ip
    transfer_port: string;    //中转port
    total_up: number;         //
    total_down: number;       //
    //vmess参数
    v: string;
    aid: string;
    scy: string;  //加密方式
    net: string; //传输协议 ws tcp kcp quic grpc等
    disguise_type: string;  //伪装类型 none http
    host: string;
    path: string;
    tls: string;
    sni: string;
}

declare interface NodeInfo {
    created_at: string;
    updated_at: string;
    id: number;
    //sspanel 参数
    node_speedlimit: number; //节点限速/Mbps
    traffic_rate: number;    //倍率
    node_type: string;       //类型 vless vmess trojan
    server: string;
    // type: string;//显示与隐藏

    //共享节点额外需要的参数
    uuid: string;
    //基础参数
    remarks: string;//别名
    address: string;
    port: number;
    node_order: number;//节点排序
    enabled: boolean;  //是否为激活节点
    //中转参数
    enable_transfer: boolean;//是否启用中转
    transfer_address: string;//中转ip
    transfer_port: number;   //中转port
    //流量
    total_up: number;
    total_down: number;

    goods: [];//多对多,关联商品
    //vmess参数
    v: string;
    scy: string;//加密方式 auto,none,chacha20-poly1305,aes-128-gcm,zero
    aid: number;//额外ID
    //vless参数
    flow: string;//流控 none,xtls-rprx-vision,xtls-rprx-vision-udp443
    encryption: string;//加密方式 none

    network: string;//传输协议 tcp,kcp,ws,h2,quic,grpc
    type: string;   //伪装类型 ws,h2：无    tcp,kcp：none，http    quic：none，srtp，utp，wechat-video，dtls，wireguard
    host: string;   //伪装域名
    path: string;   //path
    mode: string;   //grpc传输模式 gun，multi

    security: string;//传输层安全类型 none,tls,reality
    sni: string;
    fp: string;
    alpn: string;
    allowInsecure: boolean;//tls 跳过证书验证
    pbk: string;
    sid: string;
    spx: string;
}

declare interface ServerStatusInfo {
    id: number;
    name: string;
    status: boolean;
    last_time: string;
    user_amount: number;
    traffic_rate: number;
    u: number;
    d: number;
}

declare interface NodeSharedInfo{
    id: number;
    created_at: string;
    updated_at: string;

    node_type: string;
    remarks: string;
    uuid: string;
    address: string;
    port: number;
    ns: string;
    tcping: number;
    ascription: string;
    enabled: boolean;
    //vmess参数
    v: string;
    scy: string;//加密方式 auto,none,chacha20-poly1305,aes-128-gcm,zero
    aid: number;
    //vless参数
    flow: string;//流控 none,xtls-rprx-vision,xtls-rprx-vision-udp443
    encryption: string;

    network: string; ////传输协议 tcp,kcp,ws,h2,quic,grpc
    type: string;  //伪装类型 ws,h2,grpc：无    tcp：none，http    kcp,quic：none，srtp，utp，wechat-video，dtls，wireguard
    host: string;
    path: string;
    mode: string; //grpc传输模式 gun，multi

    security: string;        //传输层安全类型 none,tls,reality
    sni: string;             //
    fp: string;              //

    alpn: string;            //tls
    allowInsecure: boolean;  //tls

    pbk: string;             //reality
    sid: string;             //reality
    spx: string;             //reality
}
