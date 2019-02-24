import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios'

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      result: {}
    };      
  }

  componentDidMount() {


    axios.get('https://7siim5bbqh.execute-api.us-east-1.amazonaws.com/Prod?format=list')
    .then(response => console.log(response))

    // const groupRootUrl = "https://7siim5bbqh.execute-api.us-east-1.amazonaws.com/Prod?format=list";
    // // const groupRootUrl = "/Prod?format=list";
    // const url = groupRootUrl;// + getFilters() + addPagination();
    // fetch(url, { mode: 'no-cors',
    //     headers: {
    //       "Content-Type": "application/json",
    //       // "Access-Control-Allow-Origin": "https://7siim5bbqh.execute-api.us-east-1.amazonaws.com",
    //       // "Content-Type": "application/x-www-form-urlencoded",
    //     }
    //   })
    //   .then(
    //     res => 
    //     res.json()
    //     )
    //   .then(
    //     (result) => {
    //       this.setState({
    //         isLoaded: true,
    //         result: result,
    //       });
    //     },
    //     // Note: it's important to handle errors here
    //     // instead of a catch() block so that we don't swallow
    //     // exceptions from actual bugs in components.
    //     (error) => {
    //       this.setState({
    //         isLoaded: false,
    //         error: error.message
    //       });
    //     }
    // )
}

  render() {
    if(this.state.isLoaded) {
      console.log(this.state.result)
      return (
        <div className="App">
          <header className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <p>
             HELLO
            </p>
            <a
              className="App-link"
              href="https://reactjs.org"
              target="_blank"
              rel="noopener noreferrer"
            >
              {/* {this.state.result} */}
            </a>
          </header>
        </div>
      );
    }
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.js</code> and save to reload.
          </p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
        </header>
      </div>
    );
  
  }
}

export default App;
