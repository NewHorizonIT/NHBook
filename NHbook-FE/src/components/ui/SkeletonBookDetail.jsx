const SkeletonBookDetail = () => {
  return (
    <div className='grid grid-cols-1 lg:grid-cols-2 gap-10 animate-pulse px-4'>
      {/* Skeleton ảnh */}
      <div className='w-full flex justify-center'>
        <div className='bg-gray-300 rounded w-[300px] h-[400px]' />
      </div>

      {/* Skeleton thông tin */}
      <div className='space-y-4'>
        <div className='h-8 bg-gray-300 rounded w-2/3'></div>
        <div className='h-6 bg-gray-300 rounded w-1/4'></div>
        <div className='h-4 bg-gray-300 rounded w-1/2'></div>
        <div className='h-4 bg-gray-300 rounded w-1/3'></div>
        <div className='h-4 bg-gray-300 rounded w-full'></div>
        <div className='h-4 bg-gray-300 rounded w-full'></div>
        <div className='h-10 bg-gray-300 rounded w-32'></div>
        <div className='h-10 bg-gray-300 rounded w-1/2'></div>
      </div>
    </div>
  );
};

export default SkeletonBookDetail;
