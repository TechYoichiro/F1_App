"use client";
import React, { useState, useEffect } from 'react';

interface Driver {
  givenName: string;
  familyName: string;
}

interface Constructor {
  name: string;
}

interface RaceResult {
  position: string;
  Driver: Driver;
  Constructor: Constructor;
  Time: { time: string };
  points: string;
}

interface QualifyingResult {
  position: string;
  Driver: Driver;
  Constructor: Constructor;
  Q1: string;
  Q2: string;
  Q3: string;
}

interface Race {
  season: string;
  round: string;
  raceName: string;
  Results: RaceResult[];
  QualifyingResults: QualifyingResult[];
}

const RaceResults: React.FC = () => {
  const [race, setRace] = useState<Race | null>(null);
  const [view, setView] = useState<'qualifying' | 'race'>('race');

  useEffect(() => {
    fetch('http://localhost:8080/result')
      .then(response => response.json())
      .then(data => setRace(data));
  }, []);

  if (!race) {
    return <div>Loading...</div>;
  }

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">{race.raceName} ({race.season} - Round {race.round})</h1>
      <div className="mb-4">
        <button
          className={`px-4 py-2 mr-2 ${view === 'race' ? 'bg-blue-500 text-white' : 'bg-gray-200'}`}
          onClick={() => setView('race')}
        >
          Race
        </button>
        <button
          className={`px-4 py-2 ${view === 'qualifying' ? 'bg-blue-500 text-white' : 'bg-gray-200'}`}
          onClick={() => setView('qualifying')}
        >
          Qualifying
        </button>
      </div>
      {view === 'race' ? (
        <table className="table-auto w-full">
          <thead>
            <tr>
              <th className="px-4 py-2">POS</th>
              <th className="px-4 py-2">DRIVER</th>
              <th className="px-4 py-2">TEAM</th>
              <th className="px-4 py-2">TIME</th>
              <th className="px-4 py-2">POINT</th>
            </tr>
          </thead>
          <tbody>
            {race.Results.map(result => (
              <tr key={result.position}>
                <td className="border px-4 py-2">{result.position}</td>
                <td className="border px-4 py-2">{result.Driver.givenName} {result.Driver.familyName}</td>
                <td className="border px-4 py-2">{result.Constructor.name}</td>
                <td className="border px-4 py-2">{result.Time.time}</td>
                <td className="border px-4 py-2">{result.points}</td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <table className="table-auto w-full">
          <thead>
            <tr>
              <th className="px-4 py-2">POS</th>
              <th className="px-4 py-2">DRIVER</th>
              <th className="px-4 py-2">TEAM</th>
              <th className="px-4 py-2">Q1</th>
              <th className="px-4 py-2">Q2</th>
              <th className="px-4 py-2">Q3</th>
            </tr>
          </thead>
          <tbody>
            {race.QualifyingResults.map(result => (
              <tr key={result.position}>
                <td className="border px-4 py-2">{result.position}</td>
                <td className="border px-4 py-2">{result.Driver.givenName} {result.Driver.familyName}</td>
                <td className="border px-4 py-2">{result.Constructor.name}</td>
                <td className="border px-4 py-2">{result.Q1}</td>
                <td className="border px-4 py-2">{result.Q2}</td>
                <td className="border px-4 py-2">{result.Q3}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
}

export default RaceResults;
