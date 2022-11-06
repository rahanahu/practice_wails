import React from 'react';
import {useState} from 'react';
import {LoadUserConfig, SelectDirectory, SaveUserConfig}  from '../wailsjs/go/userconfig/UserConfig';


function ProjectSelector(){
    const [dirText, setDirText] = useState("Before initialize");
    const updateDirText = (e: any) => {
        updateNonEmptyDir(e.target.value);
    }
    const updateNonEmptyDir = (e: string) =>{
        if (e != ""){
            setDirText(e);
        }
    }


    const selectUserDir = () => {
        SelectDirectory().then(updateNonEmptyDir).catch(alert);
    };


    React.useEffect(() => {
        LoadUserConfig().then(function(result){
            updateNonEmptyDir(result.directory.toString());
        })
      }, []);
    return (
        <div>
            <div id="ProjectSelector" className="input-box">
                <input id="Dir" className="input" value={dirText} onChange={updateDirText} autoComplete="off" name="input" type="text"/>
                <button className="btn" onClick={selectUserDir} >Select Directory</button>
            </div>
        </div>
    )
}

export default ProjectSelector