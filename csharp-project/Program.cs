using System.Runtime.InteropServices;

namespace GoFunctionExample
{
    class Program
    {
        // Import the Go function for Helm list
        [DllImport("mylib", EntryPoint = "HelmList", CallingConvention = CallingConvention.Cdecl)]
        public static extern void HelmList();

        static void Main(string[] args)
        {
            Console.WriteLine("Calling Go HelmList function...");
            HelmList();
        }
    }
}
