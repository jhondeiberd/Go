public  class MercedesDealer extends AbstractFactory
    {
        @Override
        public  CAR createCAR()
        {
            return new MercedesCAR();
        }


        @Override
        public  TRUCK createTRUCK()
        {
            return new MercedesTRUCK();
        }


    }