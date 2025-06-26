import React from 'react';

const Table = ({ columns, rows }) => {
  return (
    <div className='overflow-x-auto'>
      <table className='table'>
        <thead>
          <tr>
            {columns.map((column, index) => (
              <th key={index}>{column.header}</th>
            ))}
          </tr>
        </thead>
        <tbody>
          {rows.map((row, index) => (
            <tr key={index}>
              {columns.map((column, colIndex) => (
                <td key={colIndex}>
                  {column.render
                    ? column.render(row)
                    : row[column.accessor] || 'N/A'}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Table;
