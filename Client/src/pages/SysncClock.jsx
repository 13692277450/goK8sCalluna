import React from 'react';
import {
  JsonHubProtocol,
  HubConnectionState,
  HubConnectionBuilder,
  HttpTransportType,
  LogLevel,
} from '@microsoft/signalr';

    

class SynscClock extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        message:'',
        hubConnection: null,
      };
    }
  
    componentDidMount() {
      const connection = new HubConnectionBuilder()
        .withUrl(process.env.REACT_APP_APIBASEURL+"realtime", {
        })
        .withAutomaticReconnect()
        .withHubProtocol(new JsonHubProtocol())
        .configureLogging(LogLevel.Information)
        .build();
 
    // Note: to keep the connection open the serverTimeout should be
    // larger than the KeepAlive value that is set on the server
    // keepAliveIntervalInMilliseconds default is 15000 and we are using default
    // serverTimeoutInMilliseconds default is 30000 and we are using 60000 set below
        connection.serverTimeoutInMilliseconds = 60000;
 
    // re-establish the connection if connection dropped
        connection.onclose(error => {
            console.assert(connection.state === HubConnectionState.Disconnected);
            console.log('Connection closed due to error. Try refreshing this page to restart the connection', error);
        });
    
        connection.onreconnecting(error => {
            console.assert(connection.state === HubConnectionState.Reconnecting);
            console.log('Connection lost due to error. Reconnecting.', error);
        });
    
        connection.onreconnected(connectionId => {
            console.assert(connection.state === HubConnectionState.Connected);
            console.log('Connection reestablished. Connected with connectionId', connectionId);
        });
        
        this.setState({ hubConnection: connection})

        this.startSignalRConnection(connection).then(()=> {
              if(connection.state === HubConnectionState.Connected) {
                connection.invoke('RequestSyncTime').then(val => {
                  console.log("Signalr get data first time:",val);
                  this.setState({ message:val })
                })
              }
        }) ;

        connection.on('receive', res => {
          console.log("SignalR get hot res:", res)
            this.setState({
              message:res
            });
        });
    }
  
    startSignalRConnection = async connection => {
      try {
          await connection.start();
          console.assert(connection.state === HubConnectionState.Connected);
          console.log('SignalR connection established');
      } catch (err) {
          console.assert(connection.state === HubConnectionState.Disconnected);
          console.error('SignalR Connection Error: ', err);
          setTimeout(() => this.startSignalRConnection(connection), 5000);
      }
    };
  
     render() {
      return (
        <div style={{width: '300px',float:'left',marginLeft:'10px'}} >
          <h4>最新同步完成时间： {this.state.message}  </h4>
        </div>
      );
    }
  }
export  default  SynscClock;