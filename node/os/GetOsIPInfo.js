const os = require('os');

/**
 * GetIP is a wrapper of nodejs os.networkInterfaces
 */
class GetIP {
    constructor() {
        this.IPInfo = FlapInfo(os.networkInterfaces());
    }
    /**
     * get ip info by family & internal
     *
     * @param {string} family  IPv4 or IPv6
     * @param {bool} internal  whether ip is a loopback or similar interface that is not remotely accessible
     */
    getIPInfo(family = 'IPv4', internal = true) {
        return this.IPInfo.find(
            item => item.family === family && item.internal === internal
        );
    }
    /**
     * show all local IP informations
     */
    getIPInfoList() {
        return this.IPInfo;
    }
    /**
     * show all network interface name
     */
    getNetworkInterfaceList() {
        return Object.keys(wanInfo);
    }
}

function FlapInfo(wanInfo) {
    let arr = [];
    Object.keys(wanInfo).forEach(item => {
        let tem = wanInfo[item].map(i => {
            i.NetworkInterface = item;
            return i;
        });
        arr = [...arr, ...tem];
    });
    return arr;
}

module.exports = GetIP;
