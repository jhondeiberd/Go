using System;
using System.Collections.Generic;
using System.Text;

namespace FactoryMethod
{
    class Circle : Shape
    {

        public void draw()
        {
            Console.WriteLine("inside Circle::draw() method");
        }
    }
}
