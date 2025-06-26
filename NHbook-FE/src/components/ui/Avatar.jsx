import clsx from 'clsx';
import React from 'react';

const Avatar = ({
  thumbnail = 'https://th.bing.com/th/id/OIP.EhKlVZLzgrF0kHjzodRdIAHaHa?r=0&rs=1&pid=ImgDetMain',
  className,
}) => {
  return (
    <div className='avatar'>
      <div
        className={clsx(
          'ring-primary ring-offset-base-100 rounded-full ring-2 ring-offset-2 overflow-hidden',
          `w-6`,
          !!className && className
        )}
      >
        <img src={thumbnail} />
      </div>
    </div>
  );
};

export default Avatar;
