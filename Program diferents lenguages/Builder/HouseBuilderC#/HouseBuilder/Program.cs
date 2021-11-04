using System;

namespace HouseBuilder
{
    class Program
    {
        
        static void Main(string[] args)
        {
            //the client chooses the model of product he wants
            
            HouseBuilder basicBuilder = new BasicHouseBuilder();

            //the client delegates the project of his house to the engineer

            CivilEngineer engineer = new CivilEngineer(basicBuilder);

            //the engineer builds the house for the client.

            engineer.constructHouse();

            //the engineer delivers the house to the client.

            Console.WriteLine("builder constructed basement : " + engineer.getHouse().basement);
            Console.WriteLine("builder constructed interior : " + engineer.getHouse().interior);
            Console.WriteLine("builder constructed roof : " + engineer.getHouse().roof);
            Console.WriteLine("builder constructed structure : " + engineer.getHouse().structure);


            Console.WriteLine("-------------------------------------------");

            HouseBuilder highEndBuilder = new HighEndHouseBuilder();
            CivilEngineer engineer2 = new CivilEngineer(highEndBuilder);

            //the engineer builds the house for the client.

            engineer2.constructHouse();

            //the engineer delivers the house to the client.

            Console.WriteLine("builder constructed basement : " + engineer2.getHouse().basement);
            Console.WriteLine("builder constructed interior : " + engineer2.getHouse().interior);
            Console.WriteLine("builder constructed roof : " + engineer2.getHouse().roof);
            Console.WriteLine("builder constructed structure : " + engineer2.getHouse().structure);


        }
    }
}
