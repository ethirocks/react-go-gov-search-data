import React, { useState, useEffect } from "react";
import axios from "axios";

// Define the structure of the data returned by the backend
interface GovernmentData {
  state: string;
  population: string;
  state_fips: string;
}

const DataDisplay: React.FC = () => {
  const [data, setData] = useState<GovernmentData[]>([]); // Expecting an array of objects
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [filter, setFilter] = useState("");

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get("http://localhost:8080/api/data");

        if (response.data.success) {
          setData(response.data.data); // Set the array of objects
        } else {
          setError("Failed to fetch data.");
        }
      } catch (err) {
        setError("Failed to fetch data from the server.");
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, []);

  // Handle filtering based on the state name
  const filteredData = data.filter((item) =>
    item.state.toLowerCase().includes(filter.toLowerCase())
  );

  if (loading) return <p>Loading...</p>;
  if (error) return <p>{error}</p>;

  return (
    <div>
      <h2>Government Data</h2>
      <input
        type="text"
        placeholder="Filter by state"
        value={filter}
        onChange={(e) => setFilter(e.target.value)}
        style={{
          padding: "8px",
          margin: "10px 0",
          width: "100%",
          maxWidth: "300px",
          border: "1px solid #ccc",
          borderRadius: "4px",
        }}
      />
      {filteredData.length > 0 ? (
        <table border={1} cellPadding={10} style={{ marginTop: "10px" }}>
          <thead>
            <tr>
              <th>State</th>
              <th>Population</th>
              <th>State FIPS</th>
            </tr>
          </thead>
          <tbody>
            {filteredData.map((item, index) => (
              <tr key={index}>
                <td>{item.state}</td>
                <td>{item.population}</td>
                <td>{item.state_fips}</td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <p>No results found.</p>
      )}
    </div>
  );
};

export default DataDisplay;
