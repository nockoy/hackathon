import TextField from '@mui/material/TextField';

export default function SearchBox() {
  return (
      <div className='SearchBox'
        style={{
          width:"80vw"
        }}>
        <TextField
          fullWidth
          id="outlined-textarea"
          label="slack-like app内を検索"
          placeholder="メッセージや返信を検索してみよう"
          size="small"
          margin="dense"
        />
      </div>  
  );
}