import React, { createContext, useState } from "react";

export const DNAContext = createContext({
  Text: null,
  setText: null,
  Username: null,
  setUsername: null,
  Disease: null,
  setDisease: null,
  Method: null,
  setMethod: null,
  cleanable: null,
  setCleanable: null,
  isLoading: null,
  setLoading: null,
  data: null,
  setData: null,
  file: null,
  setFile: null,
  error: null,
  setError: null,
  searchData: null,
  setSearchData: null,
  searchRes: null,
  setSearchRes: null,
  diseaseList: null,
  setDiseaseList: null,
});

export function DNAProvider({ children }) {
  const [Text, setText] = useState("");
  const [Username, setUsername] = useState("");
  const [Disease, setDisease] = useState("");
  const [Method, setMethod] = useState("");
  const [cleanable, setCleanable] = useState(null);
  const [isLoading, setLoading] = useState(null);
  const [file, setFile] = useState(null);
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  const [searchData, setSearchData] = useState(null);
  const [searchRes, setSearchRes] = useState(null);
  const [diseaseList, setDiseaseList] = useState(null);

  return (
    <DNAContext.Provider
      value={{
        Text,
        setText,
        Username,
        setUsername,
        Disease,
        setDisease,
        Method,
        setMethod,
        cleanable,
        setCleanable,
        isLoading,
        setLoading,
        file,
        setFile,
        data,
        setData,
        error,
        setError,
        searchData,
        setSearchData,
        searchRes,
        setSearchRes,
        diseaseList,
        setDiseaseList,
      }}
    >
      {children}
    </DNAContext.Provider>
  );
}
