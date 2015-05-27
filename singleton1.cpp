class HeapDataHolders
{
public:
        static HeapDataHolders& GetHeaps()
        {
                static HeapDataHolders heaps;
                return heaps;
        }

        ~HeapDataHolders()
        {
        }

private:
        explicit HeapDataHolders() {}

private:
        HeapDataHolders(const HeapDataHolders&);
        void operator=(const HeapDataHolders&);

};

