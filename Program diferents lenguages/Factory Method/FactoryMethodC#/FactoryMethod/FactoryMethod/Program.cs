using System;

namespace FactoryMethod
{
    class Program
    {
        static void Main(string[] args)
        {

            ShapeFactory shapeFactory = new ShapeFactory();

            //get a copy of class Circle and invoke the method draw() on it.
            Shape shape1 =   shapeFactory.getShape("CIRCLE");
            shape1.draw();


            Shape shape2 = shapeFactory.getShape("RECTANGLE");
            shape2.draw();

            Shape shape3 = shapeFactory.getShape("SQUARE");
            shape3.draw();


            


        }
    }
}
