using System.Linq;
using System.Net;
using System.Net.NetworkInformation;

namespace XWIN.Utilities
{
    public static class NetworkUtility
    {
        public static bool IsPortActive(int port)
        {
            IPGlobalProperties ipGlobalProperties = IPGlobalProperties.GetIPGlobalProperties();
            IPEndPoint[] activeEndPoints = ipGlobalProperties.GetActiveTcpListeners();

            return activeEndPoints.Any(endPoint => endPoint.Port == port);
        }
    }
}