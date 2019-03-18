import React, { Component } from 'react';
import './App.css';
import Operations from './Operations'

class App extends Component {

  title = () => {
    return(<div id="title">
      <h1>Will you please excuse me</h1>
      <p>When there is no excuse for not having an excuse</p>
  </div>)
  }

  operations = () => {
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
            <td>{item.message}</td>
          </tr>
          ))}
        </table>
      : 
        <p>Loading ...</p>
      }
    </div>
  )
  }

  api = () => {
    return (
      <div>
        <h2>Api Functionalities</h2>
        <h3>Content-type</h3>
        <ul>
          <li>text/html</li>
          <p>This is the default response content type for when a request is made to the page</p>
          <li>application/json</li>
          <p>in the format</p>
          <code>{'{"to":"to", "from":"from", "memo":"excuse goes here"}'}</code>
        </ul>
        <p>Format</p>
        <ul>
          <li>slack</li>
            <p>This return the response in a format that can be interpreted by slack to automatically intergrate with there format</p>
            <p>It will always repond as json</p>
            <code>{'{"text":"excuse goes here"}'}</code>
          <li>list</li>
            <p>This returns a list of all the possible commands and how they respond</p>
            <p>It will always repond as json</p>
        </ul>
      </div>
    )
  }

  inegrations = () => {
    return (
    <div>
      <h2>Integrations</h2>
      <h3>Slack</h3>
      <p>IN PROGRESS needs to be released</p>
      <p>you can add your own slack bot all you need to do is specify url:</p>
      <code>{this.url + "?format=slack"}</code>
    </div>
    )
  }


  goals = () => {
    return(
      <div>
        <h2>future goals</h2>
        <p>Discord</p>
        <p>split list apart into sperate function</p>
        <p>use diferent sub modules for excuses</p>
        <p>add searching ability to this website</p>
        <p>params that can take in multi value</p>
      </div>
    )
  }

  about = () => {
    return(     
      <div>
        <h2>about me</h2>
        <p>I was at work and couldnt focus on my assigned work so i decided to spice life up a little and write a new service in go
        (you why not avoid doing real work when you can just make excuses). While at the same time a friend was asking me for help
        so natuary I wanted to send him one of the the nice and friendlt messages from foass.com. But while those are fun they didn't
        capture my desire to have an excuse about not helping him or not doing my work or just why i dont want to adult some times. So 
        I decided to use my slacking off to make sure I always could have an excuse handy.</p>
      </div>
    )
  }

  render() {

    // title "Will you please excuse me"
    // intro (one liner) "When there is no excuse for not having an excuse"
    // api functionalities 
    //    Content-type
    //    Format
    //    
    // integrations (aka slack)
    // operations
    //     iternate over the list endpoint
    // future goals
        // split list apart into sperate function
        // use diferent sub modules for excuses
        // add searching ability to this website
        // params that can take in multi value
    // about me
        // I was at work and couldnt focus on my assigned work so i decided to spice life up a little and write a new service in go
        // (you why not avoid doing real work when you can just make excuses). While at the same time a friend was asking me for help
        // so natuary I wanted to send him one of the the nice and friendlt messages from foass.com. But while those are fun they didn't
        // capture my desire to have an excuse about not helping him or not doing my work or just why i dont want to adult some times. So 
        // I decided to use my slacking off to make sure I always could have an excuse handy.


    return(
      <div>
        {this.title()}
        {/* table of contents maybe */}
        <Operations />
        {this.api()}
        {this.inegrations()}
        {this.goals()}
        {this.about()}
      </div>)
  
  }
}

export default App;
