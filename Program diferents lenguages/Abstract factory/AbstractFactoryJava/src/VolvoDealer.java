 public class VolvoDealer extends AbstractFactory
    {
        @Override
        public  CAR createCAR()
        {
            return new VolvoCAR();
        }
        @Override
        public  TRUCK createTRUCK()
        {
            return new VolvoTRUCK();
        }

    }