import {IAceEditorProps} from 'react-ace/src/ace';
import AceEditor from 'react-ace';
import {Label} from '../Label/Label';
import 'ace-builds/src-noconflict/mode-sh';
import 'ace-builds/src-noconflict/theme-github';

export type CodeEditorProps = Omit<IAceEditorProps, 'theme' | 'setOptions'> & {
  lang: 'sh';
  label: string;
};

export const CodeEditor = ({
                             lang,
                             label,
                             ...editorProps
                           }: CodeEditorProps) => {
  return (
    <div className={'my-2'}>
      <Label>
        <span>{label}</span>
        <AceEditor
          className={'rounded-lg w-full h-64'}
          mode={lang}
          theme="github"
          editorProps={{$blockScrolling: true}}
          showPrintMargin={true}
          showGutter={true}
          highlightActiveLine={true}
          height={'200px'}
          width={'100%'}
          setOptions={{
            enableBasicAutocompletion: true,
            enableLiveAutocompletion: true,
            enableSnippets: true,
            showLineNumbers: true,
            tabSize: 4,
          }}
          {...editorProps}
        />
      </Label>
    </div>
  );
};
