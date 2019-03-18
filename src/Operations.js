import React, { Component } from 'react';

class Operations extends Component {

  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      result: {},
    };      
    this.url = "https://api.will-you-excuse.me/"

  }

  

  componentDidMount() {
    fetch(this.url + "?format=list", 
        { method: "GET",
          headers: {
            "Content-Type": "application/json",
            // "Content-Type": "application/x-www-form-urlencoded",
          }
        })
        .then(res => res.json())
        .then(
          (result) => {
            this.setState({
              isLoaded: true,
              result: result,
            });
          },
          // Note: it's important to handle errors here
          // instead of a catch() block so that we don't swallow
          // exceptions from actual bugs in components.
          (error) => {
            this.setState({
              isLoaded: false,
              error: error.message
            });
          }
        )
  }


    render(){
        return(
            <div>
                <h2>Operations</h2>
                {this.state.isLoaded?
                <table>
                    <tr>
                    <th>Key</th>
                    <th>Message</th>
                    </tr>
                    {this.state.result.map(item => (
                    <tr>
                    <td>{item.key}</td>
                    <td>{item.doc}</td>
                    <td><a href={this.url + "?key=" + item.key}>Go To</a></td>
                
                    </tr>
                    ))}
                </table>
                : 
                <p>Loading ...</p>
                }
            </div> 
        )
    }
}
export default Operations;
