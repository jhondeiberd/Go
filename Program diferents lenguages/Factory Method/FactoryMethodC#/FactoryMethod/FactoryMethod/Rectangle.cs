using System;
using System.Collections.Generic;
using System.Text;

namespace FactoryMethod
{
    class Rectangle : Shape
    {
        public void draw()
        {
            Console.WriteLine("inside Rectangle::draw() mehtod");
        }
    }
}
