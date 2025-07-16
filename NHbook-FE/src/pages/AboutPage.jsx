import React from 'react';

const AboutPage = () => {
  return (
    <div className='bg-transparent min-h-screen py-16 px-4 sm:px-8 lg:px-24 text-content'>
      <div className='max-w-6xl mx-auto grid md:grid-cols-2 gap-12 items-center'>
        {/* Left content */}
        <div>
          <h1 className='text-4xl sm:text-5xl font-extrabold text-primary mb-6'>
            Về NHBook
          </h1>

          <p className='text-lg leading-relaxed mb-4'>
            <span className='font-semibold'>NHBook</span> là nền tảng bán sách
            trực tuyến hiện đại, cung cấp hàng ngàn đầu sách chất lượng về công
            nghệ, kinh tế, văn học và đời sống.
          </p>

          <p className='text-lg leading-relaxed mb-4'>
            Chúng tôi hướng đến việc xây dựng cộng đồng yêu sách với trải nghiệm
            mua sắm thông minh, thuận tiện và đầy cảm hứng. Mỗi cuốn sách là một
            hành trình và NHBook luôn đồng hành cùng bạn.
          </p>

          <p className='text-lg leading-relaxed mb-4'>
            Thành lập từ năm 2024, NHBook đã phục vụ hàng nghìn khách hàng trên
            khắp Việt Nam với sự tận tâm, minh bạch và chuyên nghiệp.
          </p>

          <div className='mt-8 space-y-2'>
            <h2 className='text-xl font-semibold text-blue-600'>
              Thông tin liên hệ
            </h2>
            <p>
              <strong>Email:</strong> support@nhbook.vn
            </p>
            <p>
              <strong>Hotline:</strong> 0901 234 567
            </p>
            <p>
              <strong>Địa chỉ:</strong> 123 Đường Sách, Quận 1, TP.HCM
            </p>
          </div>
        </div>

        {/* Right image */}
        <div className='w-full'>
          <img
            src='https://images.unsplash.com/photo-1524985069026-dd778a71c7b4?auto=format&fit=crop&w=800&q=80'
            alt='NHBook Library'
            className='rounded-2xl shadow-xl w-full object-cover'
          />
        </div>
      </div>
    </div>
  );
};

export default AboutPage;
