using System;
using System.Runtime.InteropServices;
using System.Text.Json;

namespace GoFunctionExample
{
    class Program
    {
        // Import the HelmList function from the shared library
        [DllImport("mylib", EntryPoint = "HelmList", CallingConvention = CallingConvention.Cdecl)]
        public static extern IntPtr HelmList();

        // Define a class for Helm release information
        class Release
        {
            public string Name { get; set; }
            public string Namespace { get; set; }
            public int Revision { get; set; }
            public string Updated { get; set; }
            public string Status { get; set; }
            public string Chart { get; set; }
            public string AppVersion { get; set; }
        }

        static void Main(string[] args)
        {
            // Call the Go function
            IntPtr resultPtr = HelmList();
            string resultJson = Marshal.PtrToStringAnsi(resultPtr) ?? string.Empty;

            Console.WriteLine("Raw JSON received from Go:");
            Console.WriteLine(resultJson);

            // Deserialize the JSON string into C# objects
            var releases = JsonSerializer.Deserialize<Release[]>(resultJson, new JsonSerializerOptions { PropertyNameCaseInsensitive = true });

            // Print the Helm releases
            Console.WriteLine("\nDeserialized Helm Releases:");
            if (releases is not null)
            {
                foreach (var release in releases)
                {
                    Console.WriteLine($"Name: {release.Name}, Namespace: {release.Namespace}, Chart: {release.Chart}, Status: {release.Status}");
                }
            }
            else
            {
                Console.WriteLine("Deserialization returned null.");
            }
        }
    }
}
